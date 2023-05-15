import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';

export type OrderDetailDocument = HydratedDocument<OrderDetail>;

@Schema({ timestamps: true })
export class OrderDetail {
  @Prop({ default: uuidv4, required: true })
  id: string;

  @Prop({ required: true })
  orderId: string;

  @Prop({ required: true })
  ticketId: string;

  @Prop({ required: true, default: now() })
  createdAt: Date;

  @Prop({ required: true, default: now() })
  updatedAt: Date;

  @Prop({ default: null })
  deletedAt: Date;
}

export const OrderDetailSchema = SchemaFactory.createForClass(OrderDetail);
