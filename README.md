# Localization utility

This is a simple utility to help with localization of strings in a project. It is designed to be used with a project that uses a simple key-value pair system for localization.

## Usage
Fill the required fields in the config.yaml file. Then run the following command to generate the localization files.
```bash
./localization
```

## Configuration file
```yaml
origin_file_extension: dart
origin_locale: es
output_locales: [en, fr, de]
output_folder: output
input_folder: test_folder
```