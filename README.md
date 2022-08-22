# Star
Display some information about a github or sourcehut repository.

# Build and Install
```sh
make
sudo make install
```

# Usage
GitHub does not allow making very many "unauthenticated" API calls and SourceHut
does not allow any. So, you'll want to go generate auth tokens for each of them.

[github](https://github.com/settings/tokens)\
[sourcehut](https://meta.sr.ht/oauth/personal-token)

Export them both as environment variables before running star:
```sh
GITHUB_AUTH_TOKEN=github_key_goes_here
SOURCEHUT_AUTH_TOKEN=sourcehut_key_goes_here
```

Then, simply run `star` with a new-line separated list of repo URLs as standard
input:
```sh
star < repo_list.txt
```


# Author
Written and maintained by Dakota Walsh.
Up-to-date sources can be found at https://git.sr.ht/~kota/star/

# License
GNU GPL version 3 only, see LICENSE.
