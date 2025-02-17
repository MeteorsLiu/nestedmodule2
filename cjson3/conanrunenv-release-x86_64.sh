script_folder="/home/runner/work/nestedmodule2/nestedmodule2/cjson3"
echo "echo Restoring environment" > "$script_folder/deactivate_conanrunenv-release-x86_64.sh"
for v in LD_LIBRARY_PATH DYLD_LIBRARY_PATH
do
    is_defined="true"
    value=$(printenv $v) || is_defined="" || true
    if [ -n "$value" ] || [ -n "$is_defined" ]
    then
        echo export "$v='$value'" >> "$script_folder/deactivate_conanrunenv-release-x86_64.sh"
    else
        echo unset $v >> "$script_folder/deactivate_conanrunenv-release-x86_64.sh"
    fi
done


export LD_LIBRARY_PATH="/home/runner/.conan2/p/b/mpfrb9428f2632fa8/p/lib:/home/runner/.conan2/p/b/gmp35b895bd4ce93/p/lib:$LD_LIBRARY_PATH"
export DYLD_LIBRARY_PATH="/home/runner/.conan2/p/b/mpfrb9428f2632fa8/p/lib:/home/runner/.conan2/p/b/gmp35b895bd4ce93/p/lib:$DYLD_LIBRARY_PATH"