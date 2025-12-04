#include <stdio.h>
#include <string.h>
#include <lmdb.h>

int main () {
    char command[50];
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

        scanf("%s", command);
    };

    return 0;
};
