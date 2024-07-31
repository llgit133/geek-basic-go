



webook 项目结构学习 后端golang + gin + gorm
// Package domain 放置领域对象      User
// Package dao 代表对数据库的操作    UserDAO
// Package repository 代表领域对象的存储 UserRepository  DTO VO
// Package service 代表领域服务 UserService



User
UserDAO
UserRepository
UserService


1.Package domain
这一层包含了应用程序的核心概念和业务逻辑。
在DDD中，这通常被称为“领域模型”。User 类型会在这里定义，它表示了业务领域中的用户实体，可能包含属性如用户名、密码、电子邮件等，以及与这些属性相关的业务规则和行为。

2.Package dao
DAO（Data Access Object）层负责提供对底层数据存储的访问。
UserDAO 是一个数据访问对象，它处理与用户数据的具体读写操作，如从数据库中读取用户记录或将新用户记录保存到数据库中。这个层通常不包含复杂的业务逻辑，而是专注于数据的持久化。

3.Package repository
Repository 层充当了领域模型和数据访问之间的桥梁。
UserRepository 提供了一个抽象的数据访问接口，使得领域层可以以一种更面向对象的方式与数据交互，而不需要关心具体的数据库操作。这使得领域层的代码更加干净和可测试。

4.Package service
Service 层实现了业务逻辑和服务，UserService 就是在这里定义的。
它可能包含各种与用户相关的业务规则和流程，比如用户注册、登录、权限检查等。服务层通常依赖于领域模型和Repository层，但并不直接与数据存储交互，而是通过调用Repository来获取或修改数据。





对比springboot

1.DAO (Data Access Object) 层: 在Go代码中，dao.NewUserDAO(db) 
    类似于Spring Boot中的Repository层，它直接与数据库交互

2.Repository 层: Go中的repository.NewUserRepository(ud) 
    在Spring Boot中通常是不必要的，因为Spring Data JPA已经提供了Repository层的实现

3.service.NewUserService(ur) 
    类似于Spring Boot中的Service层，它包含了业务逻辑。

4.Controller/Handler 层: web.NewUserHandler(us) 
    类似于Spring Boot中的Controller层，它处理HTTP请求。
    在Spring Boot中，这通常是一个带有@RestController或@Controller注解的类，

5.路由注册: 在Go代码中，hdl.RegisterRoutes(server) 相当于在Spring Boot中配置路由。
    在Spring Boot中，路由是通过Controller的方法上的@RequestMapping、@GetMapping、@PostMapping等注解自动注册的，






webook-fe 项目结构学习 前端Vue项目