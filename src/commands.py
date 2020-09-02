from os import system, path, makedirs

class Commands:
    def __init__(self):
        pass

    @staticmethod
    def add_user(self, username, password,expiration, userPath, dataPublic):
        homeUser = path.join(userPath, username)
        homeUserUpload = path.join(userPath, username, "upload")
        publicUser = path.join(userPath, username, "upload")
        self.ensure_dir(homeUserUpload)
        self.ensure_dir(publicUser)
        system("useradd --groups sftpgroup --shell=/bin/false -d {} {}".format(homeUser, username))
        system('echo -e "{}\\n{}" | passwd {}'.format(password, password, username))
        Commands.permissionUser(username=username,
                                homeUser=homeUser,
                                homeUserUpload=homeUserUpload)
        system("mount --bind {} {}".format(dataPublic, publicUser))

    def permissionUser(self, username,homeUser, homeUserUpload):
        system("chown root {}".format(homeUser))
        system("chmod go-w {}".format(homeUser))
        system("chown -R {}:sftpgroup {}".format(username, homeUserUpload))
        system("chmod ug+rwX {}".format(homeUserUpload))

    def ensure_dir(file_path):
        directory = path.dirname(file_path)
        if not path.exists(directory):
            makedirs(directory)


