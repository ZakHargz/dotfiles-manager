return {
  "nvim-telescope/telescope.nvim",
  branch = "0.1.x",
  dependencies = {
    "nvim-lua/plenary.nvim",
    "nvim-tree/nvim-web-devicons",
    { "nvim-telescope/telescope-fzf-native.nvim", build = "make" },
  },
  config = function()
    local telescope = require("telescope")
    local actions = require("telescope.actions")

    telescope.setup({
      defaults = {
        mappings = {
          i = {
            ["<C-k>"] = actions.move_selection_previous,
            ["<C-j>"] = actions.move_selection_next,
            ["<C-q>"] = actions.send_selected_to_qflist + actions.open_qflist,
          },
        },
        file_ignore_patterns = { "node_modules", "git", ".next" },
      },
      extensions = {
        file_browser = {
          initial_mode = "normal",
          sorting_strategy = "ascending",
        },
      },
    })

    telescope.load_extension("fzf")
    telescope.load_extension("file_browser")

    local keymap = vim.keymap

    keymap.set("n", "<leader>ff", "<cmd>Telescope find_files<CR>", { desc = "Find files in cwd" })
    keymap.set("n", "<leader>fg", "<cmd>Telescope git_files<CR>", { desc = "Find files in cwd" })
    keymap.set("n", "<leader>b", "<cmd>Telescope buffers<CR>", { desc = "Find files in cwd" })
    keymap.set("n", "<leader>fb", "<cmd>Telescope file_browser<CR>", { desc = "Find files in cwd" })
    keymap.set(
      "n",
      "<leader>fe",
      "<cmd>Telescope file_browser path=%:p:h select_buffer=true<CR>",
      { desc = "Find files in cwd" }
    )
  end,
}
