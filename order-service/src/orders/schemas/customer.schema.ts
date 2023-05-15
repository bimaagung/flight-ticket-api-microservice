import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';

export type CustomerDocument = HydratedDocument<Customer>;

@Schema({ timestamps: true })
export class Customer {
  @Prop({ required: true })
  id: string;

  @Prop({ required: true })
  name: string;

  @Prop({ required: true })
  gender: number;

  @Prop({ required: true })
  address: string;

  @Prop({ required: true })
  email: string;

  @Prop({ type: Date, required: true, default: now() })
  createdAt: Date;

  @Prop({ type: Date, required: true, default: now() })
  updatedAt: Date;

  @Prop({ type: Date, default: null })
  deletedAt: Date;
}

export const CustomerSchema = SchemaFactory.createForClass(Customer);
