import os
import shutil
import re

def fix_contest_structure():
    contests_dir = 'contests'
    if not os.path.exists(contests_dir):
        print(f"Error: '{contests_dir}' directory not found.")
        return

    for contest_id in os.listdir(contests_dir):
        contest_path = os.path.join(contests_dir, contest_id)
        if not os.path.isdir(contest_path):
            continue

        for tag in os.listdir(contest_path):
            tag_path = os.path.join(contest_path, tag)
            if not os.path.isdir(tag_path):
                continue

            for problem_folder in os.listdir(tag_path):
                problem_folder_path = os.path.join(tag_path, problem_folder)
                if not os.path.isdir(problem_folder_path):
                    continue

                go_file_name = f"{problem_folder}.go"
                go_file_path = os.path.join(problem_folder_path, go_file_name)

                if os.path.exists(go_file_path):
                    destination_path = os.path.join(tag_path, go_file_name)
                    print(f"Moving '{go_file_path}' to '{destination_path}'")
                    shutil.move(go_file_path, destination_path)
                    print(f"Removing redundant directory '{problem_folder_path}'")
                    shutil.rmtree(problem_folder_path)

    print("\nRestructuring complete.")

if __name__ == '__main__':
    fix_contest_structure()