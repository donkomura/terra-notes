if [ -d "environments/$GITHUB_REF_HEAD/" ]; then
  cd environments/$GITHUB_REF_HEAD
  terraform apply -auto-approve
else
  echo "***************************** SKIPPING APPLYING *******************************"
  echo "Branch '$BRANCH_NAME' does not represent an oficial environment."
  echo "*******************************************************************************"
fi
