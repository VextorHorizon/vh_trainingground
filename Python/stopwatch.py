#!/usr/bin/env python3
# Modern Standby‑style Stopwatch (PySide6)
# — Minimal, high‑contrast, keyboard‑first, always‑on‑top option.
# Hotkeys: Space = Start/Pause, R = Reset, L = Lap, H = Hide/Show UI,
#          T = Toggle Always on Top, F = Fullscreen, Esc = Exit Fullscreen

from __future__ import annotations

import json
import os
import sys
import time
from pathlib import Path
from typing import List, Tuple

from PySide6.QtCore import Qt, QTimer, QSize
from PySide6.QtGui import QFont, QFontDatabase, QAction, QCloseEvent
from PySide6.QtWidgets import (
    QApplication,
    QMainWindow,
    QWidget,
    QLabel,
    QPushButton,
    QHBoxLayout,
    QVBoxLayout,
    QListWidget,
    QCheckBox,
)


SETTINGS_PATH = Path.home() / ".modern_stopwatch.json"


def format_elapsed(ms: int) -> str:
    total_seconds = ms // 1000
    hundredths = (ms % 1000) // 10
    hours = total_seconds // 3600
    minutes = (total_seconds % 3600) // 60
    seconds = total_seconds % 60
    if hours > 0:
        return f"{hours:02d}:{minutes:02d}:{seconds:02d}.{hundredths:02d}"
    else:
        return f"{minutes:02d}:{seconds:02d}.{hundredths:02d}"


class Stopwatch:
    def __init__(self) -> None:
        self._running = False
        self._start_ref = 0.0
        self._acc_ms = 0
        self._last_lap_ms = 0
        self.laps: List[Tuple[int, int]] = []  # (total_ms, delta_ms)

    def start(self) -> None:
        if not self._running:
            self._running = True
            self._start_ref = time.perf_counter()

    def pause(self) -> None:
        if self._running:
            now = time.perf_counter()
            self._acc_ms += int((now - self._start_ref) * 1000)
            self._running = False

    def reset(self) -> None:
        self._running = False
        self._start_ref = 0.0
        self._acc_ms = 0
        self._last_lap_ms = 0
        self.laps.clear()

    def toggle(self) -> None:
        if self._running:
            self.pause()
        else:
            self.start()

    @property
    def running(self) -> bool:
        return self._running

    @property
    def elapsed_ms(self) -> int:
        if self._running:
            return self._acc_ms + int((time.perf_counter() - self._start_ref) * 1000)
        return self._acc_ms

    def lap(self) -> Tuple[int, int]:
        total = self.elapsed_ms
        delta = total - self._last_lap_ms
        self._last_lap_ms = total
        self.laps.append((total, delta))
        return total, delta


class StopwatchWindow(QMainWindow):
    def __init__(self) -> None:
        super().__init__()
        self.setWindowTitle("Standby Stopwatch")
        self.setMinimumSize(QSize(520, 200))

        # Try to enable numerical tabular figures if available
        QFontDatabase.addApplicationFontFromData(b"")  # no-op; placeholder to keep API consistent

        self.sw = Stopwatch()
        self.timer = QTimer(self)
        self.timer.setInterval(10)  # 10 ms for hundredths
        self.timer.timeout.connect(self._tick)

        self._build_ui()
        self._load_settings()
        self._apply_on_top(self.top_chk.isChecked())

        # Start UI refresh timer (always on; stopwatch state decides content)
        self.timer.start()

    def _build_ui(self) -> None:
        central = QWidget(self)
        self.setCentralWidget(central)

        # Time label — scales with window size
        self.time_label = QLabel("00:00.00", self)
        self.time_label.setObjectName("timeLabel")
        self.time_label.setAlignment(Qt.AlignCenter)
        self._set_time_font()

        # Controls
        self.start_btn = QPushButton("Start")
        self.start_btn.setObjectName("start")
        self.start_btn.clicked.connect(self._toggle_start)

        self.lap_btn = QPushButton("Lap")
        self.lap_btn.setObjectName("lap")
        self.lap_btn.clicked.connect(self._add_lap)
        self.lap_btn.setEnabled(False)

        self.reset_btn = QPushButton("Reset")
        self.reset_btn.setObjectName("reset")
        self.reset_btn.clicked.connect(self._reset)
        self.reset_btn.setEnabled(False)

        self.top_chk = QCheckBox("On Top")
        self.top_chk.stateChanged.connect(lambda _: self._apply_on_top(self.top_chk.isChecked()))
        self.top_chk.setChecked(True)

        ctrl_row = QHBoxLayout()
        ctrl_row.setSpacing(10)
        ctrl_row.addStretch(1)
        ctrl_row.addWidget(self.start_btn)
        ctrl_row.addWidget(self.lap_btn)
        ctrl_row.addWidget(self.reset_btn)
        ctrl_row.addSpacing(12)
        ctrl_row.addWidget(self.top_chk)
        ctrl_row.addStretch(1)

        # Lap list (minimal)
        self.lap_list = QListWidget()
        self.lap_list.setVisible(True)
        self.lap_list.setMinimumHeight(80)

        # Layout
        outer = QVBoxLayout(central)
        outer.setContentsMargins(22, 22, 22, 22)
        outer.setSpacing(14)
        outer.addWidget(self.time_label, 1)
        outer.addLayout(ctrl_row)
        outer.addWidget(self.lap_list, 1)

        # Styling (dark, minimal)
        self.setStyleSheet(
            """
            QMainWindow, QWidget { background-color: #0b0b0f; color: #e9ecef; }
            QLabel#timeLabel {
                color: #f2f4f8;
            }
            QPushButton {
                background-color: rgba(255,255,255,0.04);
                border: 1px solid rgba(255,255,255,0.08);
                border-radius: 12px;
                padding: 10px 16px;
                font-weight: 600;
            }
            QPushButton:hover { border-color: rgba(255,255,255,0.18); }
            QPushButton:pressed { background-color: rgba(255,255,255,0.08); }
            QPushButton#start { color: #0eead6; border-color: rgba(14,234,214,0.3); }
            QPushButton#reset { color: #ff6b6b; border-color: rgba(255,107,107,0.3); }
            QPushButton#lap   { color: #a0b3ff; border-color: rgba(160,179,255,0.3); }
            QListWidget {
                background: transparent;
                border: 1px solid rgba(255,255,255,0.06);
                border-radius: 12px;
                padding: 6px;
            }
            QCheckBox { spacing: 8px; }
            """
        )

    # ----- Core actions -----
    def _tick(self) -> None:
        ms = self.sw.elapsed_ms
        self.time_label.setText(format_elapsed(ms))
        self._auto_size_font()

    def _toggle_start(self) -> None:
        self.sw.toggle()
        running = self.sw.running
        self.start_btn.setText("Pause" if running else "Start")
        self.lap_btn.setEnabled(running)
        # Allow reset whenever not running (and non‑zero)
        self.reset_btn.setEnabled((not running) and (self.sw.elapsed_ms > 0))

    def _add_lap(self) -> None:
        total, delta = self.sw.lap()
        self.lap_list.insertItem(0, f"Lap {len(self.sw.laps):02d}  {format_elapsed(total)}   (+{format_elapsed(delta)})")

    def _reset(self) -> None:
        self.sw.reset()
        self.time_label.setText("00:00.00")
        self.lap_list.clear()
        self.start_btn.setText("Start")
        self.lap_btn.setEnabled(False)
        self.reset_btn.setEnabled(False)

    def _apply_on_top(self, on: bool) -> None:
        flags = self.windowFlags()
        if on:
            flags |= Qt.WindowStaysOnTopHint
        else:
            flags &= ~Qt.WindowStaysOnTopHint
        self.setWindowFlags(flags)
        self.show()  # re‑show to apply flags

    # ----- Persistence -----
    def _load_settings(self) -> None:
        if SETTINGS_PATH.exists():
            try:
                data = json.loads(SETTINGS_PATH.read_text(encoding="utf-8"))
                geom = data.get("geometry")
                if geom:
                    self.restoreGeometry(bytes.fromhex(geom))
                self.top_chk.setChecked(bool(data.get("on_top", True)))
                hidden = bool(data.get("hide_ui", False))
                self._set_ui_visible(not hidden)
            except Exception:
                pass

    def _save_settings(self) -> None:
        data = {
            "geometry": self.saveGeometry().hex(),
            "on_top": self.top_chk.isChecked(),
            "hide_ui": not self.lap_list.isVisible(),
        }
        try:
            SETTINGS_PATH.write_text(json.dumps(data, ensure_ascii=False, indent=2), encoding="utf-8")
        except Exception:
            pass

    # ----- Fonts & sizing -----
    def _set_time_font(self) -> None:
        # Prefer a clean sans with tabular figures if available
        f = QFont()
        f.setFamily("Segoe UI")  # reasonable default on Windows; Qt will fallback on others
        f.setStyleHint(QFont.SansSerif)
        f.setWeight(QFont.DemiBold)
        f.setLetterSpacing(QFont.PercentageSpacing, 102)
        self.time_label.setFont(f)

    def _auto_size_font(self) -> None:
        # Scale font to fill available width/height while leaving margins
        label = self.time_label
        rect = label.contentsRect()
        if rect.width() <= 0 or rect.height() <= 0:
            return
        text = label.text()
        # Binary search point size
        low, high = 8, 500
        while low < high:
            mid = (low + high + 1) // 2
            f = label.font()
            f.setPointSize(mid)
            label.setFont(f)
            if label.fontMetrics().horizontalAdvance(text) <= rect.width() and label.fontMetrics().height() * 1.15 <= rect.height():
                low = mid
            else:
                high = mid - 1
        f = label.font()
        f.setPointSize(low)
        label.setFont(f)

    def resizeEvent(self, event) -> None:  # type: ignore[override]
        super().resizeEvent(event)
        self._auto_size_font()

    # ----- UI visibility & hotkeys -----
    def _set_ui_visible(self, visible: bool) -> None:
        self.lap_list.setVisible(visible)
        self.start_btn.setVisible(visible)
        self.lap_btn.setVisible(visible)
        self.reset_btn.setVisible(visible)
        self.top_chk.setVisible(visible)

    def keyPressEvent(self, event) -> None:  # type: ignore[override]
        key = event.key()
        mods = event.modifiers()
        if key == Qt.Key_Space:
            self._toggle_start()
            return
        if key in (Qt.Key_R,):
            self._reset()
            return
        if key in (Qt.Key_L,):
            if self.sw.running:
                self._add_lap()
            return
        if key in (Qt.Key_H,):
            self._set_ui_visible(not self.lap_list.isVisible())
            return
        if key in (Qt.Key_T,):
            self.top_chk.setChecked(not self.top_chk.isChecked())
            return
        if key in (Qt.Key_F,):
            if self.isFullScreen():
                self.showNormal()
            else:
                self.showFullScreen()
            return
        if key in (Qt.Key_Escape,):
            if self.isFullScreen():
                self.showNormal()
                return
            super().keyPressEvent(event)
        else:
            super().keyPressEvent(event)

    def closeEvent(self, event: QCloseEvent) -> None:  # type: ignore[override]
        self._save_settings()
        super().closeEvent(event)


def main():
    app = QApplication(sys.argv)
    # HiDPI (Qt 6). If enum/method not available on your Qt build, just skip.
    try:
        QApplication.setHighDpiScaleFactorRoundingPolicy(
            Qt.HighDpiScaleFactorRoundingPolicy.PassThrough
        )
    except Exception:
        pass

    win = StopwatchWindow()
    win.resize(900, 300)
    # Place roughly left side — user can move/snap as desired
    win.move(60, 120)
    win.show()
    sys.exit(app.exec())


if __name__ == "__main__":
    main()
