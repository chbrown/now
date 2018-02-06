OPTIONSETS=('' '-d' '-m' '-ms' '-ns'
            '-local' '-local -d' '-local -m' '-local -ms' '-local -ns'
            '-epoch' '-epoch -ms' '-epoch -ns')

for OPTIONS in "${OPTIONSETS[@]}"; do
  printf '| %10s | `%s` |\n' "$OPTIONS" $(NOW=2016-11-20T18:46:29.882112334Z now $OPTIONS)
done
