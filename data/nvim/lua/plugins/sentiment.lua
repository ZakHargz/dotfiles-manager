return {
  "utilyre/sentiment.nvim",
  event = { "InsertCharPre", "InsertEnter" },
  config = function()
    local sentiment = require("sentiment")

    sentiment.setup()
  end
}
