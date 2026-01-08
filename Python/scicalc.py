#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
SciCalc REPL — Advanced scientific calculator (no GUI)
Usage: python scicalc.py

Features:
- Angle mode (rad/deg) toggled with :rad / :deg
- Safe expression evaluation (sandboxed). ^ is accepted as exponent and converted to **
- Scientific functions, constants, numeric calculus, statistics, matrices, unit/base conversions
- Define variables and reuse the last answer in `ans`

Notes:
- Calculus / equation solving are numeric (not symbolic)
- User functions for calculus/roots use lambda, e.g. deriv(lambda x: x**2, 2.0)
"""

import math as _m
import cmath as _cm
import statistics as _stats
from decimal import Decimal, getcontext
import random as _rnd
from typing import List, Callable, Any

# -------------------- Angle mode --------------------
ANGLE_MODE = {"mode": "rad"}  # "rad" or "deg"

def _to_rad(x):
    # works for complex too
    return x if ANGLE_MODE["mode"] == "rad" else (x * _m.pi / 180)

def set_deg(): ANGLE_MODE["mode"] = "deg"
def set_rad(): ANGLE_MODE["mode"] = "rad"
def is_deg():  return ANGLE_MODE["mode"] == "deg"

# Trig functions that respect the angle mode (use cmath to support complex)
def sin(x):  return _cm.sin(_to_rad(x))
def cos(x):  return _cm.cos(_to_rad(x))
def tan(x):  return _cm.tan(_to_rad(x))
def asin(x):
    y = _cm.asin(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi
def acos(x):
    y = _cm.acos(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi
def atan(x):
    y = _cm.atan(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi

# Hyperbolic functions (kept consistent with original behavior)
def sinh(x): return _cm.sinh(_to_rad(x)) if is_deg() else _cm.sinh(x)
def cosh(x): return _cm.cosh(_to_rad(x)) if is_deg() else _cm.cosh(x)
def tanh(x): return _cm.tanh(_to_rad(x)) if is_deg() else _cm.tanh(x)
def asinh(x):
    y = _cm.asinh(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi
def acosh(x):
    y = _cm.acosh(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi
def atanh(x):
    y = _cm.atanh(x)
    return y if ANGLE_MODE["mode"]=="rad" else y*180/_m.pi

# -------------------- Constants --------------------
pi   = _m.pi
e    = _m.e
tau  = _m.tau
phi  = (1 + 5**0.5) / 2  # golden ratio

# Basic physical constants (SI)
c        = 299_792_458.0             # m/s
G        = 6.67430e-11               # m^3·kg^-1·s^-2
h        = 6.62607015e-34            # J·s
hbar     = h / (2*_m.pi)             # J·s
kB       = 1.380649e-23              # J/K
NA       = 6.02214076e23             # 1/mol
e_charge = 1.602176634e-19           # C
g        = 9.80665                   # m/s^2 (standard)

# -------------------- Basic math helpers --------------------
def ln(x):    return _cm.log(x)
def log10(x): return _cm.log10(x)
def log(x, base=_m.e):
    return _cm.log(x)/_cm.log(base)

def sqrt(x):  return _cm.sqrt(x)

def fact(n: int) -> int:
    if n < 0 or int(n)!=n:
        raise ValueError("factorial accepts non-negative integers only")
    return _m.factorial(int(n))

def nCr(n, r):
    n, r = int(n), int(r)
    if r<0 or n<0 or r>n: return 0
    return fact(n)//(fact(r)*fact(n-r))

def nPr(n, r):
    n, r = int(n), int(r)
    if r<0 or n<0 or r>n: return 0
    return fact(n)//fact(n-r)

# -------------------- Numerical calculus --------------------
def deriv(f: Callable[[float], float], x: float, h: float=1e-6, order:int=1):
    """Central finite difference for 1st/2nd derivative"""
    if order == 1:
        return (f(x+h) - f(x-h)) / (2*h)
    elif order == 2:
        return (f(x+h) - 2*f(x) + f(x-h)) / (h*h)
    else:
        raise ValueError("only order 1 or 2 supported")

def integrate(f: Callable[[float], float], a: float, b: float, n:int=1000):
    """Simpson's rule (n will be made even if it's odd)"""
    if n % 2 == 1: n += 1
    h = (b - a) / n
    s = f(a) + f(b)
    for k in range(1, n):
        x = a + k*h
        s += 4*f(x) if k % 2 == 1 else 2*f(x)
    return s * h/3

# -------------------- Root finding / equation solving --------------------
def newton(f, x0, df=None, tol=1e-12, maxiter=100):
    """Newton-Raphson (if df not provided, numeric derivative is used)"""
    x = x0
    for _ in range(maxiter):
        fx = f(x)
        dfx = df(x) if df else deriv(f, x)
        if dfx == 0:
            raise ZeroDivisionError("df(x) = 0")
        step = fx/dfx
        x_new = x - step
        if abs(x_new - x) < tol:
            return x_new
        x = x_new
    return x

def secant(f, x0, x1, tol=1e-12, maxiter=200):
    x_prev, x = x0, x1
    for _ in range(maxiter):
        f0, f1 = f(x_prev), f(x)
        if (f1 - f0) == 0:
            return x
        x2 = x - f1*(x - x_prev)/(f1 - f0)
        if abs(x2 - x) < tol:
            return x2
        x_prev, x = x, x2
    return x

def bisection(f, a, b, tol=1e-12, maxiter=200):
    fa, fb = f(a), f(b)
    if fa*fb > 0:
        raise ValueError("f(a) and f(b) must have opposite signs")
    for _ in range(maxiter):
        m = (a+b)/2
        fm = f(m)
        if abs(fm) < tol or abs(b-a) < tol:
            return m
        if fa*fm < 0:
            b, fb = m, fm
        else:
            a, fa = m, fm
    return (a+b)/2

def quad(a,b,c):
    """Solve ax^2 + bx + c = 0 (returns two roots, may be complex)"""
    D = b*b - 4*a*c
    r1 = (-b + _cm.sqrt(D))/(2*a)
    r2 = (-b - _cm.sqrt(D))/(2*a)
    return (r1, r2)

# -------------------- Statistics --------------------
def mean(xs):   return _stats.mean(xs)
def median(xs): return _stats.median(xs)
def var(xs):    return _stats.pvariance(xs)
def stdev(xs):  return _stats.pstdev(xs)

# -------------------- Matrices (simple pure-python implementations) --------------------
def mat(rows: List[List[float]]): return [list(r) for r in rows]

def m_shape(A):
    return (len(A), len(A[0]) if A else 0)

def m_T(A):
    return [list(row) for row in zip(*A)]

def m_add(A,B):
    ra, ca = m_shape(A); rb, cb = m_shape(B)
    if (ra,ca) != (rb,cb): raise ValueError("matrix sizes must match")
    return [[A[i][j]+B[i][j] for j in range(ca)] for i in range(ra)]

def m_mul(A,B):
    ra, ca = m_shape(A); rb, cb = m_shape(B)
    if ca != rb: raise ValueError("inner dimensions do not match for multiplication")
    BT = m_T(B)
    return [[sum(a*b for a,b in zip(row, col)) for col in BT] for row in A]

def m_det(A):
    n, m = m_shape(A)
    if n != m: raise ValueError("matrix must be square")
    if n == 1: return A[0][0]
    if n == 2: return A[0][0]*A[1][1] - A[0][1]*A[1][0]
    # Laplace expansion (slow but concise)
    s = 0
    for j in range(n):
        minor = [row[:j] + row[j+1:] for row in A[1:]]
        s += ((-1)**j) * A[0][j] * m_det(minor)
    return s

def m_inv(A):
    n, m = m_shape(A)
    if n != m: raise ValueError("matrix must be square")
    # Gauss-Jordan
    M = [row + [0]*n for row in A]
    for i in range(n): M[i][n+i] = 1
    # forward elimination / normalization
    for i in range(n):
        pivot = M[i][i]
        if pivot == 0:
            # try to swap rows
            for k in range(i+1, n):
                if M[k][i] != 0:
                    M[i], M[k] = M[k], M[i]; pivot = M[i][i]; break
        if pivot == 0:
            raise ValueError("matrix is singular (det=0)")
        # normalize row i
        factor = pivot
        M[i] = [x/factor for x in M[i]]
        # eliminate other rows
        for r in range(n):
            if r == i: continue
            factor = M[r][i]
            M[r] = [rv - factor*iv for rv, iv in zip(M[r], M[i])]
    return [row[n:] for row in M]

# -------------------- Base conversions --------------------
def to_base(n: int, base: int) -> str:
    if base < 2 or base > 36: raise ValueError("base must be 2..36")
    neg = n < 0
    n = abs(int(n))
    digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    if n == 0: return "0"
    out = []
    while n:
        out.append(digits[n % base])
        n //= base
    s = "".join(reversed(out))
    return "-" + s if neg else s

def from_base(s: str, base: int) -> int:
    return int(s, base)

def to_bin(n): return bin(int(n))
def to_oct(n): return oct(int(n))
def to_hex(n): return hex(int(n))

# -------------------- Unit conversions --------------------
_length = {
    "m":1.0, "km":1e3, "cm":1e-2, "mm":1e-3, "um":1e-6, "nm":1e-9,
    "mi":1609.344, "yd":0.9144, "ft":0.3048, "in":0.0254
}
_mass = {
    "kg":1.0, "g":1e-3, "mg":1e-6, "lb":0.45359237, "oz":0.028349523125
}
_time = {"s":1.0, "ms":1e-3, "us":1e-6, "ns":1e-9, "min":60, "h":3600, "day":86400}

def convert(value: float, from_u: str, to_u: str) -> float:
    if from_u in _length and to_u in _length:
        return value * _length[from_u] / _length[to_u]
    if from_u in _mass and to_u in _mass:
        return value * _mass[from_u] / _mass[to_u]
    if from_u in _time and to_u in _time:
        return value * _time[from_u] / _time[to_u]
    raise ValueError("unsupported or mismatched units")

def tconvert(value: float, from_u: str, to_u: str) -> float:
    f, t = from_u.upper(), to_u.upper()
    if f == t: return value
    # C, F, K conversions
    if f == "C" and t == "K": return value + 273.15
    if f == "K" and t == "C": return value - 273.15
    if f == "C" and t == "F": return value*9/5 + 32
    if f == "F" and t == "C": return (value-32)*5/9
    if f == "K" and t == "F": return (value-273.15)*9/5 + 32
    if f == "F" and t == "K": return (value-32)*5/9 + 273.15
    raise ValueError("temperature units must be C/F/K")

# -------------------- Random utilities --------------------
def seed(x=None): _rnd.seed(x)
def rand(a=0.0, b=1.0): return _rnd.uniform(a,b)
def randint(a, b): return _rnd.randint(int(a), int(b))
def normal(mu=0.0, sigma=1.0): return _rnd.gauss(mu, sigma)

# -------------------- Sandbox Eval --------------------
SAFE_GLOBALS = {
    # constants
    "pi":pi, "e":e, "tau":tau, "phi":phi,
    "c":c, "G":G, "h":h, "hbar":hbar, "kB":kB, "NA":NA, "e_charge":e_charge, "g":g,
    # math functions
    "sin":sin, "cos":cos, "tan":tan, "asin":asin, "acos":acos, "atan":atan,
    "sinh":sinh, "cosh":cosh, "tanh":tanh, "asinh":asinh, "acosh":acosh, "atanh":atanh,
    "ln":ln, "log10":log10, "log":log, "sqrt":sqrt,
    "fact":fact, "nCr":nCr, "nPr":nPr,
    "abs":abs, "round":round, "min":min, "max":max,
    "int":int, "float":float, "complex":complex,
    # calculus / solvers
    "deriv":deriv, "integrate":integrate,
    "newton":newton, "secant":secant, "bisection":bisection, "quad":quad,
    # statistics
    "mean":mean, "median":median, "var":var, "stdev":stdev,
    # matrices
    "mat":mat, "m_add":m_add, "m_mul":m_mul, "m_det":m_det, "m_inv":m_inv, "m_T":m_T,
    # base conversions
    "to_base":to_base, "from_base":from_base,
    "to_bin":to_bin, "to_oct":to_oct, "to_hex":to_hex,
    # units
    "convert":convert, "tconvert":tconvert,
    # randomness
    "seed":seed, "rand":rand, "randint":randint, "normal":normal,
    # angle mode
    "set_deg":set_deg, "set_rad":set_rad,
    # standard libs (exposed intentionally when needed)
    "cmath":_cm, "math":_m,
}

def _preprocess(expr: str) -> str:
    # allow ^ as exponent
    return expr.replace("^", "**")

def safe_eval(expr: str, env: dict) -> Any:
    code = _preprocess(expr)
    return eval(code, {"__builtins__": {}}, env)

# -------------------- REPL --------------------
BANNER = """\
SciCalc (DEG/RAD, calculus, matrices, stats, units) — type :help for commands
Examples:
  sin(30)       # if DEG mode is active this yields 0.5
  set_deg()     # toggle using :deg / :rad commands
  deriv(lambda x: x**3, 2)           # derivative at x=2
  integrate(lambda x: x**2, 0, 1)    # integral 0..1
  newton(lambda x: x**2-2, 1.0)      # find sqrt(2)
  m_inv(mat([[1,2],[3,4]]))          # matrix inverse 2x2
  convert(1500, "m", "km")           # = 1.5
"""

HELP = """\
REPL commands:
  :help      show this help
  :vars      list user variables
  :deg       set angle mode to degrees
  :rad       set angle mode to radians
  :quit      exit the program

Tips:
- Use ^ for exponentiation (e.g. 2^10)
- Create variables like x = 3.14
- Last answer is stored in ans
- User functions as lambdas work, e.g. newton(lambda x: x^3-1, 0.5)
"""

def repl():
    print(BANNER)
    env = dict(SAFE_GLOBALS)
    env.update({"ans": 0})
    while True:
        try:
            s = input("⟩ ").strip()
        except (EOFError, KeyboardInterrupt):
            print("\nbye.")
            break
        if not s:
            continue
        if s.startswith(":"):
            cmd = s[1:].strip().lower()
            if cmd == "help":
                print(HELP)
            elif cmd == "vars":
                keys = sorted(k for k in env.keys() if k not in SAFE_GLOBALS)
                # show only user-defined variables
                for k in keys:
                    print(f"{k} = {env[k]}")
            elif cmd == "deg":
                set_deg(); print("Angle mode: DEG")
            elif cmd == "rad":
                set_rad(); print("Angle mode: RAD")
            elif cmd in ("quit","exit","q"):
                print("bye.")
                break
            else:
                print("Unknown command, see :help")
            continue

        # assignment?
        if "=" in s and not s.lstrip().startswith(("==", "!=","<=",">=")):
            # simple pattern a = expression
            name, expr = s.split("=", 1)
            name = name.strip()
            expr = expr.strip()
            if not name.isidentifier():
                print("invalid variable name")
                continue
            try:
                val = safe_eval(expr, env)
                env[name] = val
                env["ans"] = val
                print(val)
            except Exception as e:
                print(f"ERROR: {e}")
            continue

        try:
            val = safe_eval(s, env)
            env["ans"] = val
            print(val)
        except Exception as e:
            print(f"ERROR: {e}")

if __name__ == "__main__":
    repl()
