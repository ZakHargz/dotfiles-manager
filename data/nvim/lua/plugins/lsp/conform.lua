return {
  "stevearc/conform.nvim",
  event = { "BufReadPre", "BufNewFile" },
  opts = {
    formatters_by_ft = {
      lua = { "stylua" },
      typescript = { "prettier", "prettierd" },
      javascript = { "prettier", "prettierd" },
      typescriptreact = { "prettier", "prettierd" },
      javascriptreact = { "prettier", "prettierd" },
      json = { "prettier", "prettierd" },
      yaml = { "prettier", "prettierd" },
      python = { "black", "isort" },
    },
    format_on_save = {
      lsp_fallback = true,
      async = false,
      timeout_ms = 500,
    },
    notify_on_error = true,
  },
}
