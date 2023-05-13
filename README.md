# toolkit

开发项目用到的工具库, 借鉴了部分开源项目的utility, 避免每次都造轮子。


## ds包
该包包含常见的数据结构以及一些函数式方法

### Slice 切片工具包
* Slice操作 SliceOp*
* Slice读取 SliceGet*
* Slice转换 SliceConvert*
* Slice包含 SliceInclude*
* Slice比较 SliceCmp*
* Slice最值 SliceMax*，SliceMin*
* Slice分组 SliceGroup*

### Map map工具包
* Map操作 MapOp*
* Map比较 MapCmp*
* Map转换 MapConvert*

### Tuple 元组包
* NewTupleN 构造元组


### Fp函数式相关
* xxIter\* <br>
  迭代集合中元素
* xxIterAllOk\* <br>
  迭代集合中元素，判断是否**所有元素都满足条件**, 返回布尔值
* xxIterFilter\* <br>
  迭代集合中元素，过滤元素
* xxIterMapInPlace\* <br>
  迭代集合中元素，本地修改，不返回变更结果
* xxIterMapCopy\* <br>
  迭代集合中元素，修改副本返回结果
* xxIterPartition\* <br>
  迭代集合中的元素，分成两批元素
* FpPartial\* <br>
  绑定函数参数
