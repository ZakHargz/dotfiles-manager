function clone() # Clone repo and install dependencies
{
	gh repo clone "$1"
	cd "$(basename "$_" .git)" || exit

	if test -f "./package.json"; then
		echo "Installing NPM dependencies.."
		npm install
		return 0
	elif test -f "./pyproject.toml"; then
		echo "Installing Poetry dependencies..."
        poetry init
		poetry install
		return 0
	elif test -f "./requirements.txt"; then
		python3 -m venv ./venv
		source venv/bin/activate
		pip install -r requirements.txt
		echo "Installing dependenciesfrom requirements.txt..."
		return 0
	elif test -f "./composer.json"; then
		echo "Installing composer dependencies..."
		composer install
		return 0
	fi
}
