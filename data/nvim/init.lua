local modules = {
  "core",
  "lazyinit"
}

for _, module in ipairs(modules) do
  require(module)
end
