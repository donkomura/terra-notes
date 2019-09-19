if [ -d "environments/$GITHUB_HEAD_REF/" ]; then
  cd environments/$GITHUB_REF_HEAD
  terraform init
else
  for dir in environments/*/
  do
    cd ${dir}
    env=${dir%*/}
    env=${env#*/}
    echo ""
    echo "*************** TERRAFORM INIT ******************"
    echo "******* At environment: ${env} ********"
    echo "*************************************************"
    terraform init || exit 1
    cd ../../
  done
fi

