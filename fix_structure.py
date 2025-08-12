import os
import shutil
import re

def fix_contest_structure():
    contests_dir = 'contests'
    if not os.path.exists(contests_dir):
        print(f"Error: '{contests_dir}' directory not found.")
        return

    for item in os.listdir(contests_dir):
        item_path = os.path.join(contests_dir, item)
        if not os.path.isdir(item_path):
            continue

        # The folder name is currently like '1945D', we need to separate '1945' and 'D'
        match = re.match(r'^(\d+)([A-Z])$', item)
        if match:
            contest_id = match.group(1)
            problem_letter = match.group(2)
            
            print(f"Processing: {item}")

            # Create the new contest directory if it doesn't exist
            new_contest_dir = os.path.join(contests_dir, contest_id)
            if not os.path.exists(new_contest_dir):
                os.makedirs(new_contest_dir)
                print(f"  Created directory: {new_contest_dir}")

            # Create the problem directory
            problem_dir = os.path.join(new_contest_dir, problem_letter)
            if not os.path.exists(problem_dir):
                os.makedirs(problem_dir)
                print(f"  Created directory: {problem_dir}")

            # Move the entire old folder to the new problem directory
            # The old folder contains the .go file
            destination_path = os.path.join(problem_dir, item)
            shutil.move(item_path, destination_path)
            print(f"  Moved '{item}' to '{problem_dir}/'")
        else:
            print(f"Skipping item with unexpected name format: {item}")

    print("\nRestructuring complete.")

if __name__ == '__main__':
    fix_contest_structure()