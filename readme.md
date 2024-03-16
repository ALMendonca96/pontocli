# Pontocli

Pontocli is a command-line interface (CLI) application for tracking work hours. It allows users to store and manage the times they worked on various days.

## Features

- **Add Work Hours**: Record the hours you worked on specific dates, including entry, lunch break, return from lunch, and exit times.
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
pontocli add --date YYYY-MM-DD --hour=HH:mm
```

To add more than one work hour, the hour parameter followed by ,

```bash
pontocli add --date YYYY-MM-DD --hour=HH:mm,HH:mm
```

To add the current date and hour, just omit the date and hour parameters

```bash
pontocli add
```

Replace YYYY-MM-DD, HH:MM, etc., with the appropriate date and time values.

### Viewing Work Hours

To view the recorded work hours for a date, use the view command:

```bash
pontocli view --date=YYYY-MM-DD
```

Replace YYYY-MM-DD with the date for which you want to view the work hours.

### Deleting Work Hours

To delete specific work hours for a date, use the delete command:

```bash
pontocli delete --date=YYYY-MM-DD --hour=HH:MM,HH:MM
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

```

```
