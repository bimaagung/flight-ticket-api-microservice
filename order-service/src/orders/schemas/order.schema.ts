import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';

export type OrderDocument = HydratedDocument<Order>;

@Schema({ timestamps: true })
export class Order {
  @Prop({ default: uuidv4, required: true })
  id: string;

  @Prop({ required: true })
  customerId: string;

  @Prop({ required: true })
  qty: number;

  @Prop({ required: true })
  amount: number;

  @Prop({ type: Date, required: true, default: now() })
  createdAt: Date;

  @Prop({ type: Date, required: true, default: now() })
  updatedAt: Date;

  @Prop({ type: Date, default: null })
  deletedAt: Date;
}

export const OrderSchema = SchemaFactory.createForClass(Order);
