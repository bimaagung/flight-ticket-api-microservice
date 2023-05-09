import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';

export type OrderDocument = HydratedDocument<Order>;

@Schema()
export class Order {
  @Prop({ required: true })
  id: string;

  @Prop({ required: true })
  userId: string;

  @Prop({ required: true })
  ticketId: string;

  @Prop({ required: true })
  qty: number;

  @Prop({ required: true })
  amount: number;

  @Prop({ required: true, default: now() })
  createdAt: Date;

  @Prop({ required: true, default: now() })
  updatedAt: Date;

  @Prop({ default: null })
  deletedAt: Date;
}

export const OrderSchema = SchemaFactory.createForClass(Order);
