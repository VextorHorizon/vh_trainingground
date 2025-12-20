class AppConfig:
    def __init___(self, name, app_type, path):
        self.name = name
        self.type = app_type
        self.path = path

    def launch(self):
        print(f"Launching {self.name} from {self.path}...")

# with open('profile.json', 'r') as profile:
#     profile_data = json.load(profile)

app1 = AppConfig("Vscode", "Editor", "/usr/bin/code")
app1.launch()