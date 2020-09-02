from classes.commands import Commands
import json

class User:
    def __init__(self, username, expirationDate, password):
        f = open("./appsettings.json")
        data = json.load(f)
        self.__username = username
        self.__expirationDate = expirationDate
        self.__password = password
        self.__pathUsers = data['usersPath']
        self.__pathPublic = data['publicPath']

    @property
    def get_username(self):
        return self.__username

    @get_username.setter
    def set_username(self, username):
        self.__username = username

    @property
    def get_password(self):
        return self.__password

    @get_password.setter
    def set_password(self, password):
        self.__password = password

    @property
    def get_usersPath(self):
        return self.__pathUsers

    @property
    def get_publicPath(self):
        return self.__pathPublic

    def create_user(self):
        Commands.add_user(
            username=self.get_username(),
            password=self.get_password(),
            expiration=7,
            userPath=self.get_usersPath(),
            dataPublic=self.get_publicPath())