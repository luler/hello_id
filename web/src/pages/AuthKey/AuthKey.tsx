import {PlusOutlined} from '@ant-design/icons';
import {
  ModalForm,
  PageContainer,
  ProFormDigit,
  ProFormText,
  ProFormTextArea,
  ProTable
} from '@ant-design/pro-components';
import {Button, Form, message} from 'antd';
import React, {useRef, useState} from 'react';
import {requestGet, requestPost} from "@/utils/requestTool";

const Index: React.FC = () => {
  const [modalOpen, setmodalOpen] = useState(false)
  const [form] = Form.useForm();
  const actionRef = useRef();
  const [loading, setLoading] = useState(false);
  const columns = [
    {
      title: 'ID',
      dataIndex: 'Id',
      search: false,
    },
    {
      title: '授权码',
      dataIndex: 'AuthKey',
    }, {
      title: '描述',
      dataIndex: 'Remark',
    },
    {
      title: '操作',
      search: false,
      render: record => {
        return <>
          <a onClick={() => {
            setmodalOpen(true)
            let data = {...record}
            form.setFieldsValue(data)
          }}>编辑</a>
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
          FixedLength: 0,
        })
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
    await requestGet('/api/getAuthKeyList', params).then(res => {
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
        title={form.getFieldValue('Id') > 0 ? '编辑授权码' : '新增授权码'}
        open={modalOpen}
        form={form}
        loading={loading}
        modalProps={{
          onCancel: () => setmodalOpen(false)
        }}
        onFinish={values => {
          setLoading(true)
          // @ts-ignore
          requestPost('/api/saveAuthKey', values).then(res => {
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
          name="AuthKey"
          label="授权码"
          placeholder="自动生成"
          disabled
        />
        <ProFormTextArea
          rules={[
            {
              required: true,
              message: "请输入",
            },
          ]}
          name="Remark"
          label="描述"
          placeholder="请输入"
        />
      </ModalForm>
    </PageContainer>
  );
};

export default Index;