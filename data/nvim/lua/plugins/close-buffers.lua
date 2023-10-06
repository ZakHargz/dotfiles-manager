return {
  "kazhala/close-buffers.nvim",
  event = "BufRead",
  config = function()
    require("close_buffers").setup({})

    local map = vim.keymap.set
    local opts = { noremap = true, silent = true }

    map("n", "<leader>x", ":BDelete this<CR>", opts)
    map("n", "<leader>X", ":BDelete other<CR>", opts)
    map("n", "<leader>Z", ":BDelete all<CR>", opts)
  end,
}
