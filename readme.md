# Pontocli

Pontocli is a command-line interface (CLI) application for tracking work hours. It allows users to store and manage the times they worked on various days.

## Features

- **Add Work Hours**: Record the hours you worked on specific dates.
- **View Work Hours**: Display the recorded work hours for a given date.
- **Delete Work Hours**: Remove specific hours from a recorded workday.

## Installation

### Prerequisites

- Go programming language (version 1.21.1)
- Git

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/ALMendonca96/pontocli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd pontocli
   ```

3. Build the application:

   ```bash
   go build -o pontocli cmd/main.go
   ```

4. Install the application (optional):

   ```bash
   go install cmd/main.go
   ```

## Usage

### Adding Work Hours

To add work hours for a specific date, use the `add` command:

```bash
pontocli add --date=YYYY-MM-DD --hour=HH:mm
```

To add more than one work hour, use the hour parameter followed by a comma:

```bash
pontocli add --date=YYYY-MM-DD --hour=HH:mm,HH:mm
```

To add the current date and hour, just omit the date and hour parameters:

```bash
pontocli add
```

Additional date options:

```bash
# Add work hours for yesterday
pontocli add --date=yesterday --hour=HH:mm
```

Replace YYYY-MM-DD, HH:MM, etc., with the appropriate date and time values.

### Viewing Work Hours

To view the recorded work hours for a date, use the view command:

```bash
# View work hours for a specific date
pontocli view --date=YYYY-MM-DD

# View work hours for yesterday
pontocli view --date=yesterday

# View work hours for the last recorded date
pontocli view --date=last
```

### Deleting Work Hours

To delete specific work hours for a date, use the delete command:

```bash
# Delete work hours for a specific date
pontocli delete --date=YYYY-MM-DD --hour=HH:MM,HH:MM

# Delete work hours for yesterday
pontocli delete --date=yesterday --hour=HH:MM,HH:MM

# Delete work hours for the last recorded date
pontocli delete --date=last --hour=HH:MM,HH:MM
```

Replace YYYY-MM-DD, HH:MM, etc., with the appropriate date and time values.

## Contributing

Contributions are welcome! If you'd like to contribute to Pontocli, please follow these steps:

Fork the repository
Create a new branch (git checkout -b feature)
Make your changes
Commit your changes (git commit -am 'Add new feature')
Push to the branch (git push origin feature)
Create a new pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
