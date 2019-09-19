if [ -d "environments/$GITHUB_REF_HEAD/" ]; then
  cd environments/$GITHUB_REF_HEAD
  terraform plan
else
  for dir in environments/*/
  do 
    cd ${dir}   
    env=${dir%*/}
    env=${env#*/}  
    echo ""
    echo "*************** TERRAFOM PLAN ******************"
    echo "******* At environment: ${env} ********"
    echo "*************************************************"
    terraform plan || exit 1
    cd ../../
  done
fi

