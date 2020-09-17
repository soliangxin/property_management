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
      <el-table-column align="center" label="车位编号" width="170">
        <template slot-scope="scope">
          {{ scope.row.code }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="车位名称">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="100" align="center">
        <template scope="scope">
          <span v-if="scope.row.status=== 0">闲置中</span>
          <span v-else-if="scope.row.status=== 1">使用中</span>
        </template>
      </el-table-column>
      <el-table-column label="所属小区编号" width="170" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_code }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="添加时间" :formatter="dateFormat" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_time }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getStallList } from '@/api/stall'

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
      getStallList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
