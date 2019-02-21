# akc
AuthorizedKeys client

## Components

- **akc** -- the client, gets executed as an `AuthorizedKeysCommand` from a server's SSH configuration. Calls out to **akcd**, passes on certain configured information (optional) as well as the obvious - the username of the connecting user.
- **akcd** -- the server, accepts the connecting username as a parameter (Example: `ec2-user`), potentially looks up a calling server's extra metadata information and responds with a list of all public keys which should have access to this user@server combination. **akcd** could use metadata such as EC2 instance tags for "matching" the users who should have access.
- **akcweb** -- Where you could manage which *AuthorizedKeys* sources to pull user keys from, and *which users* keys to pull. And potentially map users into *Groups*, and *Groups* could map onto *EC2 Tag:Value* combinations. Read: "Users in group *Abbott-admins* can access all nodes with *ec2tag:Service:Abbott*.

## Concepts

- *AuthorizedKeys sources* - this could be Github, Gitlab, LDAP etc.
- 
