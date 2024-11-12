### Step 1. Publish the code on GitHub

1. **Initialize a new git repository**:

    ```bash
    git init
    git add .
    git commit -m "Initial commit of in-mem-cache"
    ```

2. **Create a repository on GitHub** (or another platform) and add it as a remote:

    ```bash
    git remote add origin https://github.com/Takeso-user/in-mem-chache.git
    ```

3. **Push the code to GitHub**:

    ```bash
    git push -u origin main
    ```

### Step 2. Install and test the package

Now your package is available for installation via `go get`. Try installing it in another project to make sure everything works:

1. Create a new project and install the package:

    ```bash
    go get github.com/Takeso-user/in-mem-chache
    ```

2. Import and use the package, following the example in `README`.

### Step 3. Update the package

If you make changes to the package and want them to be available to others, simply commit the changes and push them to GitHub:

```bash
git add .
git commit -m "Update cache implementation"
git push