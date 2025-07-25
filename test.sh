git clone https://github.com/maheshz09/boilerplates.git
cd boilerplates
# Make parent folders
mkdir ansible kubernetes terraform bash docker other scripts

# Move things
mv Ansible-basics-to-Advance ansible/basics-to-advance
mv ansible-proj ansible/proj
mv ansible_role_handling ansible/error-handling
mv ansible/ansible-role ansible/role
mv End-to-End-Kubernetes-DevSecOps-Tetris-Project kubernetes/end-to-end-devsecops-tetris
mv K8-nodes kubernetes/nodes
mv kubernetes-practice kubernetes/practice
mv tf-proj terraform/tf-proj
mv bash-scripting bash/scripting
mv bash/* bash/general-scripts/
mv docker full-stack-blogging-app docker/full-stack-blogging-app
mv react-app other/react-app
mv shorturl-golang other/shorturl-golang
mv testing other/testing
mv monitoring other/monitoring

# Move scripts into script folder
mv *.sh *.yaml *.yml scripts/

git add .
git commit -m "Restructure repo into grouped folders"
git push origin master

