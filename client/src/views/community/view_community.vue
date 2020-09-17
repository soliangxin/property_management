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
      <el-table-column label="小区编号" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.community_num }}
        </template>
      </el-table-column>

      <el-table-column label="小区名称" width="160" align="center">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column label="坐落地址">
        <template slot-scope="scope">
          <span>{{ scope.row.address }}</span>
        </template>
      </el-table-column>
      <el-table-column label="开发商名称" width="180" align="center">
        <template slot-scope="scope">
          {{ scope.row.developer }}
        </template>
      </el-table-column>
      <el-table-column label="占地面积" width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.area }}
        </template>
      </el-table-column>
      <el-table-column label="总栋数" width="80" align="center">
        <template slot-scope="scope">
          {{ scope.row.total_building }}
        </template>
      </el-table-column>
      <el-table-column label="总户数" width="90" align="center">
        <template slot-scope="scope">
          {{ scope.row.total_owner }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="修建时间" :formatter="dateFormat" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.build_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="竣工时间" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.build_time }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getCommunityList } from '@/api/community'

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
      getCommunityList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    },
    // 时间格式化
    dateFormat: function(row, column) {
      var date = row[column.property]
      return date.format('yyyy-MM-dd')
    }
  }
}
</script>
