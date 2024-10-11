# Sun CLI

Sun CLI is a command-line tool designed for quickly fetching weather information.

## Available Commands

### 1. `init`

This command initializes the Sun CLI by setting your API key and default location. You will be prompted to enter your API key and the location for which you want to receive weather updates.
**Usage:**

```bash
sun init
```

### 2. `config`

Configure the Sun CLI. This command will display the current configuration values such as the API key, location, and unit of measure.

**Usage:**

```bash
sun config
```

**Available Commands:**

- `set`

  Set a configuration value. This command will set the value of a specific configuration key. The available keys are:

  - `apiKey`: Your API key from the weatherapi.com website.
  - `location`: The location for which you want to receive weather updates. This can be a city or zipcode.
  - `unit`: The unit of measure for the temperature. Can be either `metric` or `imperial`.

  **Usage:**

  ```bash
  sun config set <key> <value>
  ```

### 3. `get`

Get weather information for the current location or for an arbitrary location. This command will display the current weather information for the specified location.

**Usage:**

```bash
sun get
```

# Requirements:

- Go 1.22.1
- weatherapi.com API key

# Installation:

1. Clone the repository:

```bash
git clone https://github.com/alexleyoung/sun.git
```

2. Change into the directory:

```bash
cd sun
```

3. Build the executable:

```bash
go build
```

4. Install the executable:

```bash
go install
```

5. Run the executable:

```bash
sun
```

# License:

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
