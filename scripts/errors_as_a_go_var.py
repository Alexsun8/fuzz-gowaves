import argparse

def read_errors(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file if line.strip()]

def generate_go_file(output_file, errors):
    with open(output_file, 'w') as go_file:
        go_file.write("package main\n\n")
        go_file.write("var standardErrors = []string{\n")
        for error in errors:
            go_file.write('\t"' + error + '",\n')
        go_file.write("}\n")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Generate Go file with standard errors.')
    parser.add_argument('-i', '--input', type=str, required=True, help='Input file path')
    parser.add_argument('-o', '--output', type=str, required=True, help='Output Go file path')
    args = parser.parse_args()

    input_file = args.input
    output_file = args.output

    errors = read_errors(input_file)
    generate_go_file(output_file, errors)
