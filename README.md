写一个 todo app.
后端 golang restful api

【数据结构】

做一个app
- 开始框架 项目时间占比
- 编码 项目时间占比
  -- 学习python 父节点时间占比
  -- 实验 父节点时间占比
  -- 学会 父节点时间占比
- 编译 app 项目时间占比
- 上线 项目时间占比

完成一道菜
- 买菜 项目时间占比
- 准备 项目时间占比
- 开始做 项目时间占比
- 吃饭 项目时间占比
- 洗碗完成。 项目时间占比

【后端需求】
- user api
    - POST /api/user/login 登陆
    request
    ```json
    {
        "username": "user_name",
        "password": "password",
    }
    ```
    response
    ```json
    {
        "token": "token",
    }
    ```
    - POST /api/user/register 注册
    request
    ```json
    {
        "username": "user_name",
        "password": "password",
    }
    ```
    response
    ```json
    {
        "token": "token",
    }
    ```
- 下面的这些 api 都需要 heander token 认证，数据都需要加一个 user_id 字段，所有操作都是针对 user_id 的。
- 获取项目列表：GET /api/projects
- 获取特定项目的任务项：GET /api/project/{project_name}/items
- 标记任务项为完成：PATCH /api/project/{project_name}/item/{item_name}/done
   request 
    ```json
    {
        "method": "done", // done, undone 动作
    }
    ```
5. 增加一个任务项：POST /api/project/{project_name}/item
    request 
    ```json
    {
        "id": "item_name/sub_item_name",
        // 父节点时间占比 %
        "parent_time": 10,
    }
    ```
6. 删除一个任务项：DELETE /api/project/{project_name}/item/{item_name}
    request 
7. 修改一个任务项：PATCH /api/project/{project_name}/item/{item_name}
    request 
    ```json
    {
        "status": "done",
        "parent_time": 10,
    }
    ```

【前端需求】
1. config.json 存储一个用户，提供登陆
2. 所有数据都用 json 存储在 golang 服务器端
3. 首页只显示大的项目进度条
4. 点击进入项目后，显示数状的任务。可以选择完成状态。