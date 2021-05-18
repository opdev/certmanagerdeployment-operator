#!/usr/bin/env bash
#
# Given two directories with identically named kubernetes manifest files,
# parse each identical pair of files and compare their resources to
# see if the files address the same kubernetes resource.
#
# Requires https://github.com/mikefarah/yq/tree/v3
#
# Usage compare_yaml_names.sh directory_one/ directory_two
#


old="${1}"
new="${2}"
red=$'\e[1;31m'
grn=$'\e[1;32m'
rst=$'\e[0m'

which yq || { echo "Expected yq to be installed. Get it here: https://github.com/mikefarah/yq/tree/v3"; exit 9 ;}

# can't find the directories, stat should give an error message
stat "${old}" 1>/dev/null || exit 1
stat "${new}" 1>/dev/null || exit 1

oldFileCount=$(find "${old}" -name "*.yaml" -type f | wc -l)
newFileCount=$(find "${new}" -name "*.yaml" -type f | wc -l)

if [[ "${oldFileCount}" -gt "${newFileCount}" ]]; then
    count="$((oldFileCount-1))"
else
    count="$((newFileCount-1))"
fi
out="FILE DIR KIND NAMESPACE NAME MATCH\n"
for i in `seq 0 ${count}` ; do
    oldFile="${old}/${i}.yaml"
    newFile="${new}/${i}.yaml"

    # handle old 
    oldType="-"
    oldNamespace="-"
    oldName="-"
    if stat "${oldFile}" &>/dev/null  ; then
        oldType=$(yq r "${oldFile}" kind)
        oldNamespace=$(yq r "${oldFile}" metadata.namespace)
        oldName=$(yq r "${oldFile}" metadata.name)
    fi
    # handle new
    newType="-"
    newNamespace="-"
    newName="-"
    if stat "${newFile}" &>/dev/null  ; then
        newType=$(yq r "${newFile}" kind)
        newNamespace=$(yq r "${newFile}" metadata.namespace)
        newName=$(yq r "${newFile}" metadata.name)
    fi
    if [[ "${oldType}${oldNamespace}${oldName}" != "${newType}${newNamespace}${newName}" ]] ; then
        clr="${red}"
        sym="✖︎"
    else 
        clr="${grn}"
        sym="✔︎"
    fi

    out="${out}${i}.yaml ${old} ${oldType} ${oldNamespace} ${oldName} ${clr}${sym}${rst}\n→ ${new} ${newType} ${newNamespace} ${newName} ${clr}${sym}${rst}\n-\n"
done
echo -e "${out}" | column -t
