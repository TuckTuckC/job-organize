#include <stdio.h>
#include <string.h>
// #include <lmdb.h>

int readFile(char *filePath) {
    char line[100];
    FILE *fptr;

    fptr = fopen(filePath, "r");

    while (fgets(line, 100, fptr)) {
        printf("%s", line);
    };
    return 0;
};

int main () {
    char command[50];
    printf("Welcome to Job Organize\nEnter a command:\n");
    scanf("%s", command);

    while(strcmp(command, "quit") != 0) {
        printf("Command: %s\n", command);

        if(strcmp(command, "getUsers") == 0) {
            printf("calling getUsers...\n");
            // getUsers();
        }

        else if(strcmp(command, "getCompanies") == 0) {
            printf("calling getCompanies...\n");
            // getUsers();
        }

        else if(strcmp(command, "getUserByID") == 0) {
            printf("calling getUserByID...\n");
            // getUsers();
        }

        else if(strcmp(command, "readFile") == 0) {
            char filePath[100];
            printf("Enter the file path: \n");
            scanf("%s", filePath);
            readFile(filePath);
        }

        scanf("%s", command);
    };

    return 0;
};
