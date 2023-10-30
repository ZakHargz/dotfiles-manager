vim.g.mapleader = " "
local keymap = vim.keymap

-- Line transform using 'Alt'
keymap.set("n", "<A-j>", ":m .+1<CR>==", { noremap = true })
keymap.set("n", "<A-k>", ":m .-2<CR>==", { noremap = true })
keymap.set("i", "<A-j>", "<Esc>:m .+1<CR>==gi", { noremap = true })
keymap.set("i", "<A-k>", "<Esc>:m .-2<CR>==gi", { noremap = true })
keymap.set("v", "<A-j>", ":m '>+1<CR>gv=gv", { noremap = true })
keymap.set("v", "<A-k>", ":m '<-2<CR>gv=gv", { noremap = true })

-- Indenting
keymap.set("v", "<", "<gv")
keymap.set("v", ">", ">gv")
