<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="房产编号" width="150">
        <template slot-scope="scope">
          {{ scope.row.house_code }}
        </template>
      </el-table-column>
      <el-table-column label="姓名" width="90" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="身份证号" width="170" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.identity_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="联系方式" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.tel }}
        </template>
      </el-table-column>
      <el-table-column label="职业" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.occupation }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="出生日期" :formatter="dateFormat" width="120">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.birth }}</span>
        </template>
      </el-table-column>
      <el-table-column label="性别" width="50" align="center">
        <template scope="scope">
          <span v-if="scope.row.gender=== 0">女</span>
          <span v-else-if="scope.row.gender=== 1">男</span>
        </template>
      </el-table-column>
      <el-table-column label="成员类型" width="90" align="center">
        <template scope="scope">
          <span v-if="scope.row.owner_type=== 1">户主</span>
          <span v-else-if="scope.row.owner_type=== 2">家庭成员</span>
          <span v-else-if="scope.row.owner_type=== 3">租户</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="入住日期" :formatter="dateFormat" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.birth }}</span>
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getTenementList } from '@/api/tenement'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getTenementList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
