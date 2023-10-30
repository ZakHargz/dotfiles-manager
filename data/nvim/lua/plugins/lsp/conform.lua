return {
	"stevearc/conform.nvim",
	event = { "BufReadPre", "BufNewFile" },
	opts = {
		formatters_by_ft = {
			lua = { "stylua" },
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
