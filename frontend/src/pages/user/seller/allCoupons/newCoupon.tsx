import { Row, Col } from 'react-bootstrap';
import { useState } from 'react';
import { useForm, SubmitHandler, useFieldArray } from 'react-hook-form';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface IFormInput {
  type: string; // 'percentage', 'fixed', 'shipping'
  name: string;
  description: string;
  discount: number;
  start_date: string;
  expire_date: string;
  tags: {
    name: string;
  }[];
}

const tagStyle = {
  borderRadius: '30px',
  background: ' var(--button_light)',
  padding: '1% 1% 1% 3%',
  color: 'white',
  margin: '5px 0 5px 5px',
  width: '100%',
};

const NewSellerCoupon = () => {
  //TODO: get the init value
  // const params = useParams();
  // const id = params.coupon_id;

  // react-hook-form things
  const { register, control, handleSubmit, watch } = useForm<IFormInput>({
    defaultValues: {
      type: 'percentage',
      name: 'Coupon',
      description: 'this is description',
      discount: 0,
      start_date: '2000-1-1',
      expire_date: '2000-1-1',
      tags: [],
    },
  });
  const { fields, append, remove } = useFieldArray({
    control,
    name: 'tags',
  });
  const OnFormOutput: SubmitHandler<IFormInput> = (data) => {
    console.log(data);
    return data;
  };
  const watchAllFields = watch();

  // tags
  const [tag, setTag] = useState('');
  const addNewTag = (event: React.KeyboardEvent<HTMLInputElement>) => {
    // this addressed the magic number: https://github.com/facebook/react/issues/14512
    if (event.keyCode === 229) return;
    if (event.key === 'Enter') {
      const input = event.currentTarget.value.trim();
      if (input !== '') {
        //TODO: api request
        append({ name: input });
        setTag('');
      }
    }
  };
  const deleteTag = (index: number) => {
    //TODO: api request
    remove(index);
  };

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <form onSubmit={handleSubmit(OnFormOutput)}>
        <Row>
          {/* left half */}
          <Col xs={12} md={5} className='goods_bgW'>
            <div className='flex-wrapper' style={{ padding: '0 8% 10% 8%' }}>
              {/* sample display */}
              <div style={{ padding: '15% 10%' }}>
                <CouponItemTemplate
                  data={{
                    id: null,
                    name: watchAllFields.name,
                    policy: watchAllFields.discount.toString(),
                    date: watchAllFields.expire_date,
                    tags: [],
                    introduction: '',
                  }}
                />
              </div>
              <span className='dark'>add more tags</span>

              {/* new tag input */}
              <input
                type='text'
                placeholder='Input tags'
                className='quantity_box'
                value={tag}
                onChange={(e) => setTag(e.target.value)}
                onKeyDown={addNewTag}
                style={{ marginBottom: '10px' }}
              />

              {/* dynamic tags fields */}
              {fields.map((field, index) => (
                <div key={field.id} style={tagStyle}>
                  <Row style={{ width: '100%' }} className='center'>
                    <Col xs={1} className='center'>
                      <FontAwesomeIcon
                        icon={faTrash}
                        className='white_word pointer'
                        onClick={() => deleteTag(index)}
                      />
                    </Col>
                    <Col>{field.name}</Col>
                  </Row>
                </div>
              ))}

              {/* delete, comfirm button */}
              <div style={{ height: '50px' }} />
              <TButton text='Delete Coupon' />
              <TButton text='Confirm Changes' onClick={handleSubmit(OnFormOutput)} />
            </div>
          </Col>

          {/* right half */}
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 0% 7% 2%' }}>
              <FormItem label='Coupon Name'>
                <input
                  type='text'
                  defaultValue={watchAllFields.name}
                  {...register('name', { required: true })}
                />
              </FormItem>

              <FormItem label='Coupon description'>
                <textarea
                  defaultValue={watchAllFields.description}
                  {...register('description', { required: true })}
                />
              </FormItem>

              <FormItem label='Method'>
                <select
                  defaultValue={watchAllFields.type}
                  {...register('type', { required: true })}
                >
                  <option value='percentage'>percentage</option>
                  <option value='fixed'>fixed</option>
                  <option value='shipping'>shipping</option>
                </select>
              </FormItem>

              <FormItem label='Discount'>
                <input
                  type='number'
                  defaultValue={watchAllFields.discount}
                  {...register('discount', { required: true })}
                />
              </FormItem>

              <FormItem label='Start Date'>
                <input
                  type='date'
                  defaultValue={watchAllFields.start_date}
                  {...register('start_date', { required: true })}
                />
              </FormItem>

              <FormItem label='Expire Date'>
                <input
                  type='date'
                  defaultValue={watchAllFields.expire_date}
                  {...register('expire_date', { required: true })}
                />
              </FormItem>
            </div>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default NewSellerCoupon;
