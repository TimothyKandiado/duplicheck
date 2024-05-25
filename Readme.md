## Duplicheck

command-line utility for checking duplicate files and deleting all but one.
can keep either most recent file or file with shortest path

manual selection also possible

# Usage

duplicheck [folder1] [folder2] [folder3] ...

example:
duplicheck /home/user/Documents /home/user/Videos

if no arguments are given it checks the present working directory recursively

# Working

`duplicheck` recursively searches through all specified folders, grouping files by their file size initially. It then computes and compares SHA1 checksums to identify which files, among those with the same size, are actual duplicates. Subsequently, it organizes these duplicates into groups based on their similar hashes.

afterwards you can select which files to keep either automatically or manually


