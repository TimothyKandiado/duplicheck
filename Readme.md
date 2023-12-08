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

duplicheck recursively searches all folders specified and groups all files with same file size
sha1 checksum is computed and compared to find which files with same size are actually duplicates
then it groups duplicate files by similar hashes

afterwards you can select which files to keep either automatically or manually
