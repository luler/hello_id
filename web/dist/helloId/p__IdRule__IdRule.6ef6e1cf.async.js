"use strict";(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[984],{29595:function(ue,T,e){e.r(T),e.d(T,{default:function(){return le}});var M=e(15009),R=e.n(M),O=e(99289),W=e.n(O),G=e(97857),I=e.n(G),H=e(5574),j=e.n(H),N=e(51042),Y=e(11774),J=e(97245),K=e(37476),Z=e(31199),Q=e(5966),d=e(1413),A=e(91),L=e(22270),i=e(67294),b=e(66758),V=e(19323),t=e(85893),U=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],X=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],w=function(r,m){var h=r.fieldProps,s=r.children,F=r.params,x=r.proFieldProps,o=r.mode,c=r.valueEnum,P=r.request,f=r.showSearch,g=r.options,p=(0,A.Z)(r,U),y=(0,i.useContext)(b.Z);return(0,t.jsx)(V.Z,(0,d.Z)((0,d.Z)({valueEnum:(0,L.h)(c),request:P,params:F,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,d.Z)({options:g,mode:o,showSearch:f,getPopupContainer:y.getPopupContainer},h),ref:m,proFieldProps:x},p),{},{children:s}))},q=i.forwardRef(function(l,r){var m=l.fieldProps,h=l.children,s=l.params,F=l.proFieldProps,x=l.mode,o=l.valueEnum,c=l.request,P=l.options,f=(0,A.Z)(l,X),g=(0,d.Z)({options:P,mode:x||"multiple",labelInValue:!0,showSearch:!0,showArrow:!1,autoClearSearchValue:!0,optionLabelProp:"label"},m),p=(0,i.useContext)(b.Z);return(0,t.jsx)(V.Z,(0,d.Z)((0,d.Z)({valueEnum:(0,L.h)(o),request:c,params:s,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,d.Z)({getPopupContainer:p.getPopupContainer},g),ref:r,proFieldProps:F},f),{},{children:h}))}),k=i.forwardRef(w),_=q,E=k;E.SearchSelect=_,E.displayName="ProFormComponent";var $=E,ee=e(1414),ae=e(15867),te=e(45360),z=e(48457),re=function(){var r=(0,i.useState)(!1),m=j()(r,2),h=m[0],s=m[1],F=ee.Z.useForm(),x=j()(F,1),o=x[0],c=(0,i.useRef)(),P=(0,i.useState)(!1),f=j()(P,2),g=f[0],p=f[1],y=[{title:"ID\u6807\u8BC6",dataIndex:"Type"},{title:"\u5F53\u524D\u9012\u589E\u503C",dataIndex:"CurrentId",search:!1},{title:"\u524D\u7F00",dataIndex:"Prefix",search:!1},{title:"\u540E\u7F00",dataIndex:"Suffix",search:!1},{title:"\u56FA\u5B9A\u957F\u5EA6",dataIndex:"FixedLength",search:!1},{title:"\u64CD\u4F5C",search:!1,render:function(u){return(0,t.jsx)(t.Fragment,{children:(0,t.jsx)("a",{onClick:function(){s(!0);var a=I()({},u);a.Prefix=a.Prefix.split(",").filter(function(v){return v!==""}),a.Suffix=a.Suffix.split(",").filter(function(v){return v!==""}),o.setFieldsValue(a)},children:"\u7F16\u8F91"})})}}],ne=function(){return[(0,t.jsx)(ae.ZP,{type:"primary",icon:(0,t.jsx)(N.Z,{}),onClick:function(){o.resetFields(),o.setFieldsValue({CurrentId:0,FixedLength:0}),s(!0)},children:"\u6DFB\u52A0"},"add")]},oe=function(){var S=W()(R()().mark(function u(n){var a;return R()().wrap(function(C){for(;;)switch(C.prev=C.next){case 0:return a={data:[],success:!0,total:0},n.Page=n.current||1,n.PageSize=n.pageSize||10,C.next=5,(0,z.P3)("/api/getIdRuleList",n).then(function(B){a.success=B.code===200,a.data=B.data.List||[],a.data=a.data.map(function(D){return D.key=D.Id,D}),a.total=B.data.Total||0});case 5:return C.abrupt("return",Promise.resolve(a));case 6:case"end":return C.stop()}},u)}));return function(n){return S.apply(this,arguments)}}();return(0,t.jsxs)(Y._z,{style:{minHeight:window.innerHeight-150},children:[(0,t.jsx)(J.Z,{columns:y,search:{labelWidth:120},request:oe,toolBarRender:ne,actionRef:c}),(0,t.jsxs)(K.Y,{width:500,title:o.getFieldValue("Id")>0?"\u7F16\u8F91ID\u89C4\u5219":"\u65B0\u589EID\u89C4\u5219",open:h,form:o,loading:g,modalProps:{onCancel:function(){return s(!1)}},onFinish:function(u){var n,a;p(!0),(0,z.dv)("/api/saveIdRule",I()(I()({},u),{},{Prefix:(n=u.Prefix)===null||n===void 0?void 0:n.join(","),Suffix:(a=u.Suffix)===null||a===void 0?void 0:a.join(",")})).then(function(v){v.code===200&&(te.ZP.success(v.message),s(!1),o.resetFields(),c.current&&c.current.reload()),p(!1)})},children:[(0,t.jsx)(Z.Z,{name:"Id",hidden:!0}),(0,t.jsx)(Q.Z,{rules:[{required:!0,message:"\u8BF7\u8F93\u5165"}],name:"Type",label:"ID\u6807\u8BC6",placeholder:"\u8BF7\u8F93\u5165"}),(0,t.jsx)(Z.Z,{rules:[{type:"number",min:0}],name:"CurrentId",label:"\u5F53\u524D\u9012\u589E\u503C",placeholder:"\u8BF7\u8F93\u5165",tooltip:"ID\u9012\u589E\u7684\u8D77\u59CB\u503C"}),(0,t.jsx)($,{name:"Prefix",mode:"tags",label:"\u524D\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)($,{name:"Suffix",mode:"tags",label:"\u540E\u7F00",placeholder:"\u8BF7\u8F93\u5165",options:[{value:"\u5E74",label:"\u5E74"},{value:"\u6708",label:"\u6708"},{value:"\u65E5",label:"\u65E5"},{value:"\u65F6",label:"\u65F6"},{value:"\u5206",label:"\u5206"},{value:"\u79D2",label:"\u79D2"}]}),(0,t.jsx)(Z.Z,{rules:[{type:"number",min:0}],name:"FixedLength",label:"\u56FA\u5B9A\u957F\u5EA6",placeholder:"\u8BF7\u8F93\u5165",tooltip:"\u8BBE\u4E3A0\u65F6\u4E3A\u4E0D\u9650\u5236\u751F\u6210\u7684ID\u7684\u957F\u5EA6"})]})]})},le=re}}]);
