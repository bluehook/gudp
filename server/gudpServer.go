// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package server

//#可信UDP服务端
// GudpServer为上层业务逻辑提供一层透明的网络服务
// GudpServer维护一个全局连接列表，内部有自己的包格式
// 发包时对上层数据包进行封装，收包时解封上层数据包
type GudpServer struct {
}
