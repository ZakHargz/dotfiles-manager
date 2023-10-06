return {
  "akinsho/bufferline.nvim",
  dependencies = { "nvim-tree/nvim-web-devicons" },
  version = "*",
  opts = {
    options = {
      always_show_bufferline = true,
      show_buffer_close_icons = true,
      mode = "buffers",
      separator_style = "slant",
      diagnostics = "nvim_lsp",
      offsets = { { filetype = "NvimTree" } }
    }
  }
}
