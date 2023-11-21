import { Col, Row } from 'react-bootstrap';

import userData from '@pages/user/seller/sellerInfo.json';
import goodsData from '@pages/discover/goodsData.json';
import TButton from '@components/TButton';
import GoodsItem from '@components/GoodsItem';

const Shop = () => {
  return (
    <Row>
      <Col xs={12} md={12}>
        <div className='user_bg center'>
          <div style={{ padding: '6% 10% 6% 10%' }}>{userData.introduction}</div>
        </div>
      </Col>
      <Col xs={12} md={3} lg={2}>
        <Row className='user_icon' style={{ padding: '0 7% 0 7%' }}>
          <Col xs={12} className='center'>
            <img src={userData.imgUrl} className='user_img' />
          </Col>
          <Col xs={12}>
            <div className='center'>
              <h4 className='title_color' style={{ padding: '10% 2% 0% 2%' }}>
                <b>{userData.name}</b>
              </h4>
            </div>
            <hr className='hr' />
            <div className='center'> Products : 13 items</div>
            <TButton text='Explore Shop' url='/sellerID/shop' />
            <TButton text='Check Coupons' url='' />
          </Col>
        </Row>
      </Col>
      <Col xs={12} md={9} ld={10} style={{ padding: '1% 5% 6% 5%' }}>
        <div className='title'>All products</div>
        <hr className='hr' />
        <Row>
          {goodsData.map((data, index) => {
            return (
              <Col xs={6} md={3} key={index}>
                <GoodsItem id={data.id} name={data.name} imgUrl={data.imgUrl} isIndex={false} />
              </Col>
            );
          })}
        </Row>
      </Col>
    </Row>
  );
};

export default Shop;
