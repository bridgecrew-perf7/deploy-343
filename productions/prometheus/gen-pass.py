# https://prometheus.io/docs/guides/basic-auth/
import getpass
import bcrypt

username = input("username: ")
username = username.strip()

password = getpass.getpass("password: ")
# print("password: ", password)

hashed_password = bcrypt.hashpw(password.encode("utf-8"), bcrypt.gensalt())

print(username + ": " + hashed_password.decode())

# web.yml
# basic_auth_users:
#    admin: $2b$12$hNf2lSsxfm0.i4a.1kVpSOVyBCfIB51VRjgBUyv6kdnyTlgWj81Ay

# prometheus --web.config.file=web.yml
