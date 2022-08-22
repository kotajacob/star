# Star
Display some information about a github or sourcehut repository.

Enter something like this into STDIN:
```
https://github.com/ange-yaghi/engine-sim      
https://github.com/aadv1k/gdoc.vim
https://github.com/anuvyklack/hydra.nvim
https://github.com/uga-rosa/cmp-dictionary
https://github.com/f3fora/cmp-spell
```

And get something like this:
```
https://github.com/ange-yaghi/engine-sim
Desc: Combustion engine simulator that generates realistic audio.
Lang: C++
Stars: 3571

https://github.com/Aadv1k/gdoc.vim
Desc: Google docs integration for vim
Lang: Python
Stars: 37
Topics: google-docs, google-docs-api, neovim, neovim-plugin, python, python-3,
python3, vim, vim-configuration, vim-plug, vim-plugin, viml, vimrc

https://github.com/anuvyklack/hydra.nvim
Desc: Create custom submodes and menus
Lang: Lua
Stars: 544
Topics: lua, neovim, neovim-plugin

https://github.com/uga-rosa/cmp-dictionary
Desc: nvim-cmp source for dictionary.
Lang: Lua
Stars: 103
Topics: lua, neovim-plugin, nvim-cmp

https://github.com/f3fora/cmp-spell
Desc: spell source for nvim-cmp based on vim's spellsuggest.
Lang: Lua
Stars: 82
Topics: nvim-cmp
```

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
