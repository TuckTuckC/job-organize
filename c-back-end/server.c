#include <stdio.h>
#include <string.h>
#include <lmdb.h>

#define GET_USERS 1

int main () {
    char command[50];
    scanf("%s", command);

    while(strcmp(command, "quit") != 0) {
        if(strcmp(command, "getUsers") == 0) {
            // getUsers();
        };
    };
    return 0;
};
