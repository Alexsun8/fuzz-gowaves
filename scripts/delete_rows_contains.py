import sys
import argparse

def read_errors(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file if line.strip()]

def filter_file(input_file, output_file='output.txt', phrases=[]):
    with open(input_file, 'r') as f_in, open(output_file, 'w') as f_out:
        for line in f_in:
            if not any(phrase in line for phrase in phrases):
                f_out.write(line)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Filter lines from input file based on given error phrases.')
    parser.add_argument('-i', '--input', type=str, required=True, help='Input file path')
    parser.add_argument('-o', '--output', type=str, default='output.txt', help='Output file path (default: output.txt)')
    parser.add_argument('-e', '--errors', type=str, default='standards_errors.txt', help='File containing error phrases (default: standards_errors.txt)')
    args = parser.parse_args()

    input_file = args.input
    output_file = args.output
    errors_file = args.errors

    standards_errors = read_errors(errors_file)

    filter_file(input_file, output_file, standards_errors)
