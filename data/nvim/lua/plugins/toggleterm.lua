return {
  "akinsho/toggleterm.nvim",
  version = "*",
  config = function()
    local toggleterm = require("toggleterm")

    toggleterm.setup({
      size = 7,
      open_mapping = [[<c-b>]],
      start_in_insert = true,
      hide_numbers = true,
      persist_size = true,
      direction = "horizontal",
      close_on_exit = true
    })
  end
}
