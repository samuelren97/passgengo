# Passgengo
A (maybe) overkill CLI tool to generate passwords

![Help image](./images/help.png)

## Basic usage

### Hashing Method
Available hashing methods
```
 -hm int
        Hashing method to use. Available methods: (default 0)
                0 -> None
                1 -> SHA256
```
> Don't worry, you will be given the cleartext password as well.  

**Output Example**
```
[INFO] => Using a hashing method, cleartext password: 4)sLu!*#9(Dz
[SUCC] => Generated password: 279ce39d5753610f9aa5f90f354dceae8cc5bee991fb745269f9ccb7b8b1fc54
```

### Length
Specifies the password length
```
-l int
        The password length. Must be between 6 and 128 characters (default 12)
```

### No special characters
Disables special characters
```
-nospecial
        Remove special characters
```

### No upper characters
Disables upper case characters
```
-noupper
        Remove upper-case characters
```

### Wizard
If you do not wish to use parameters to generate the password, you may use the `wizard` switch
```
-wizard
        Generate a password using the wizard
```