---
title: Bash Cheat Sheat
---

For enforcing of bash shell in scripting use shebang at the top of file `#!/usr/bin/env bash`. For posix compatibility use conditional syntax as `[ ... ]` and dictionaries/arrays are not supported in it.

## String Quotes
```bash
name="jane"
echo "hi $name" # hi jane
echo 'hi $name' # hi $name
```

## Conditionals
**Conditions &rarr;**
```bash
[[ -z STRING ]]	        # Empty string
[[ -n STRING ]]	        # Not empty string
[[ STRING == STRING ]]	# Equal
[[ STRING != STRING ]]	# Not Equal
[[ NUM -eq NUM ]]	    # Equal
[[ NUM -ne NUM ]]	    # Not equal
[[ NUM -lt NUM ]]	    # Less than
[[ NUM -le NUM ]]	    # Less than or equal
[[ NUM -gt NUM ]]	    # Greater than
[[ NUM -ge NUM ]]	    # Greater than or equal
[[ STRING =~ STRING ]]	# Regexp
(( NUM < NUM )) 	    # Numeric conditions
```

**File Conditions &rarr;**
```bash
[[ -e FILE ]]	        # Exists
[[ -r FILE ]]	        # Readable
[[ -h FILE ]]	        # Symlink
[[ -d FILE ]]	        # Directory
[[ -w FILE ]]	        # Writable
[[ -s FILE ]]	        # Size is > 0 bytes
[[ -f FILE ]]	        # File
[[ -x FILE ]]	        # Executable
[[ FILE1 -nt FILE2 ]]	# 1 is more recent than 2
[[ FILE1 -ot FILE2 ]]	# 2 is more recent than 1
[[ FILE1 -ef FILE2 ]]	# Same files
```

## Arrays
```bash
Fruits=('Apple' 'Banana' 'Orange')
Fruits[0]="Apple"
Fruits[1]="Banana"
Fruits[2]="Orange"

# Operations
Fruits=("${Fruits[@]}" "Watermelon")    # Push
Fruits+=('Watermelon')                  # Also Push
Fruits=( "${Fruits[@]/Ap*/}" )          # Remove by regex match
unset Fruits[2]                         # Remove one item
Fruits=("${Fruits[@]}")                 # Duplicate
Fruits=("${Fruits[@]}" "${Veggies[@]}") # Concatenate
words=($(< datafile))                   # From file (split by IFS)

# Working with arrays
echo "${Fruits[0]}"           # Element #0
echo "${Fruits[-1]}"          # Last element
echo "${Fruits[@]}"           # All elements, space-separated
echo "${#Fruits[@]}"          # Number of elements
echo "${#Fruits}"             # String length of the 1st element
echo "${#Fruits[3]}"          # String length of the Nth element
echo "${Fruits[@]:3:2}"       # Range (from position 3, length 2)
echo "${!Fruits[@]}"          # Keys of all elements, space-separated

# Ieration
for i in "${arrayName[@]}"; do
  echo "$i"
done
```

## Directories 
Also called as associative array.

```bash
# Define
declare -A sounds
sounds[dog] = "bark"
sounds[cow] = "moo"

# Working with dictionaries
echo "${sounds[dog]}" # Dog's sound
echo "${sounds[@]}"   # All values
echo "${!sounds[@]}"  # All keys
echo "${#sounds[@]}"  # Number of elements
unset sounds[dog]     # Delete dog

# Iteration
for val in "${sounds[@]}"; do
  echo "$val"
done

# Iteration over keys
for key in "${!sounds[@]}"; do
  echo "$key"
done
```
