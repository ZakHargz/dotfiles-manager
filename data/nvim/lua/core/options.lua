local opt = vim.opt

opt.number = true
opt.relativenumber = false

opt.tabstop = 2
opt.shiftwidth = 2
opt.softtabstop = 2
opt.expandtab = true
opt.autoindent = true

opt.wrap = false

opt.ignorecase = true
opt.smartcase = true

opt.termguicolors = true
opt.signcolumn = "yes"

opt.updatetime = 500
opt.scrolloff = 8
opt.cursorline = true

opt.backspace = "indent,eol,start"

opt.clipboard:append("unnamedplus")
