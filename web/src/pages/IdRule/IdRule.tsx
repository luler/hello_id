import {PlusOutlined} from '@ant-design/icons';
import {ModalForm, PageContainer, ProFormDigit, ProFormSelect, ProFormText, ProTable} from '@ant-design/pro-components';
import {Button, Divider, Form, Input, message, Modal, Popconfirm} from 'antd';
import React, {useRef, useState} from 'react';
import {requestGet, requestPost} from "@/utils/requestTool";

const Index: React.FC = () => {
  const [modalOpen, setmodalOpen] = useState(false)
  const [modelTitle, setModelTitle] = useState('新增ID规则')
  const [form] = Form.useForm();
  const actionRef = useRef();
  const [loading, setLoading] = useState(false);
  const columns = [
    {
      title: 'ID标识',
      dataIndex: 'Type',
    },
    {
      title: '当前递增值',
      dataIndex: 'CurrentId',
      search: false,
    },
    {
      title: '前缀',
      dataIndex: 'Prefix',
      search: false,
    },
    {
      title: '后缀',
      dataIndex: 'Suffix',
      search: false,
    },
    {
      title: '递增值最小长度',
      dataIndex: 'MinLength',
      search: false,
    },
    {
      title: '操作',
      search: false,
      render: record => {
        return <>
          <a onClick={() => {
            requestGet('/api/createId', {Type: record.Type}).then(res => {
              if (res.code === 200) {
                Modal.info({
                  closable: true,
                  icon: null,
                  footer: null,
                  title: '标识"' + record.Type + '"成功生成如下ID',
                  content: <div style={{paddingBottom: 20, paddingTop: 10,}}>
                    <Input value={res.data[0]} disabled/>
                  </div>,
                })
                setTimeout(() => {
                  actionRef.current.reload()
                }, 1500)
              }
            })
          }}>生成</a>
          <Divider type="vertical"/>
          <a onClick={() => {
            setModelTitle('编辑ID规则')
            setmodalOpen(true)
            let data = {...record}
            data.Prefix = data.Prefix.split(',').filter((item: string) => item !== '')
            data.Suffix = data.Suffix.split(',').filter((item: string) => item !== '')
            form.setFieldsValue(data)
          }}>编辑</a>
          <Divider type="vertical"/>
          <Popconfirm
            title='您确定要删除吗？'
            description='删除后，数据将无法恢复，请慎重！'
            onConfirm={e => {
              requestPost('/api/delIdRule', {Ids: [record.Id]}).then(res => {
                if (res.code === 200) {
                  message.success(res.message)
                  actionRef.current.reload()
                }
              })
            }}
          >
            <a style={{color: 'red',}}>删除</a>
          </Popconfirm>
        </>
      }
    },
  ];
  const renderToolBar = () => {
    return [
      <Button key="add" type="primary" icon={<PlusOutlined/>} onClick={() => {
        form.resetFields()
        form.setFieldsValue({
          CurrentId: 0,
          MinLength: 0,
        })
        setModelTitle("新增ID规则")
        setmodalOpen(true)
      }}>
        添加
      </Button>,
    ];
  };
  const getData = async (params: {}) => {
    let value = {
      data: [],
      success: true,
      total: 0,
    }
    params.Page = params.current || 1
    params.PageSize = params.pageSize || 10
    await requestGet('/api/getIdRuleList', params).then(res => {
      value.success = res.code === 200 ? true : false
      value.data = res.data.List || []
      value.data = value.data.map(item => {
        item.key = item.Id
        return item
      })
      value.total = res.data.Total || 0
    })
    return Promise.resolve(value)
  }
  return (
    <PageContainer style={{minHeight: window.innerHeight - 150}}>
      <ProTable
        columns={columns}
        search={{
          labelWidth: 120,
        }}
        request={getData}
        toolBarRender={renderToolBar}
        actionRef={actionRef}
      />
      <ModalForm
        width={500}
        title={modelTitle}
        open={modalOpen}
        form={form}
        loading={loading}
        modalProps={{
          onCancel: () => setmodalOpen(false)
        }}
        onFinish={values => {
          setLoading(true)
          // @ts-ignore
          requestPost('/api/saveIdRule', {
            ...values,
            Prefix: values.Prefix?.join(','),
            Suffix: values.Suffix?.join(','),
          }).then(res => {
            if (res.code === 200) {
              message.success(res.message)
              setmodalOpen(false)
              form.resetFields()
              actionRef.current && actionRef.current.reload()
            }
            setLoading(false)
          })
        }}
      >
        <ProFormDigit
          name="Id"
          hidden={true}
        />
        <ProFormText
          rules={[
            {
              required: true,
              message: "请输入",
            },
          ]}
          name="Type"
          label="ID标识"
          placeholder="请输入"
        />
        <ProFormDigit
          rules={[
            {
              type: 'number',
              min: 0,
            },
          ]}
          name="CurrentId"
          label="当前递增值"
          placeholder="请输入"
          tooltip="ID递增的起始值"
        />
        <ProFormSelect
          name="Prefix"
          mode="tags"
          label="前缀"
          placeholder="请输入"
          options={[
            {value: '年', label: '年'},
            {value: '月', label: '月'},
            {value: '日', label: '日'},
            {value: '时', label: '时'},
            {value: '分', label: '分'},
            {value: '秒', label: '秒'},
          ]}
        />
        <ProFormSelect
          name="Suffix"
          mode="tags"
          label="后缀"
          placeholder="请输入"
          options={[
            {value: '年', label: '年'},
            {value: '月', label: '月'},
            {value: '日', label: '日'},
            {value: '时', label: '时'},
            {value: '分', label: '分'},
            {value: '秒', label: '秒'},
          ]}
        />
        <ProFormDigit
          rules={[
            {
              type: 'number',
              min: 0,
            },
          ]}
          name="MinLength"
          label="递增值最小长度"
          placeholder="请输入"
          tooltip="设为0时为不限制生成的ID的长度"
        />
      </ModalForm>
    </PageContainer>
  );
};

export default Index;
