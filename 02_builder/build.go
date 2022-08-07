package main

// ServiceProfile 服务档案，其中服务ID唯一标识一个服务实例，一种服务类型可以有多个服务实例
type ServiceProfile struct {
	Id       string           // 服务ID
	Region   *Region          // 服务所属region
	Priority int              // 服务优先级，范围0～100，值越低，优先级越高
	Load     int              // 服务负载，负载越高表示服务处理的业务压力越大
}

// Region 值对象，每个服务都唯一属于一个Region
type Region struct {
	Id      string
	Name    string
	Country string
}

// Endpoint 值对象，其中ip和port属性为不可变，如果需要变更，需要整对象替换
type Endpoint struct {
	ip   string
	port int
}

// 1: 为ServiceProfile定义一个Builder对象
type serviceProfileBuild struct {
	//  将ServiceProfile作为Builder的成员属性
	profile *ServiceProfile
}

// 2 定义构建ServiceProfile的方法
func (s *serviceProfileBuild) WithId(id string) *serviceProfileBuild {
	s.profile.Id = id
	//返回Builder接收者指针，支持链式调用
	return s
}

func (s *serviceProfileBuild) WithRegion(regionId, regionName, regionCountry string) *serviceProfileBuild {
	s.profile.Region = &Region{Id: regionId, Name: regionName, Country: regionCountry}
	return s
}

func (s *serviceProfileBuild) WithPriority(priority int) *serviceProfileBuild {
	s.profile.Priority = priority
	return s
}

func (s *serviceProfileBuild) WithLoad(load int) *serviceProfileBuild {
	s.profile.Load = load
	return s
}

// 3: 定义Build方法，在链式调用的最后调用，返回构建好的ServiceProfile
func (s *serviceProfileBuild) Build() *ServiceProfile {
	return s.profile
}

// 4: 定义一个实例化Builder对象的工厂方法
func NewServiceProfileBuilder() *serviceProfileBuild {
	return &serviceProfileBuild{profile: &ServiceProfile{}}
}