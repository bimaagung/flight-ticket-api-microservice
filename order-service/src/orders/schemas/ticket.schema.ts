import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';

export type TicketDocument = HydratedDocument<Ticket>;

@Schema({ timestamps: true })
export class Ticket {
  @Prop({ required: true })
  id: string;

  @Prop({ required: true })
  trackId: string;

  @Prop({ required: true })
  airplaneId: string;

  @Prop({ type: Date, required: true })
  date: Date;

  @Prop({ type: Date, required: true })
  time: Date;

  @Prop({ required: true })
  price: number;

  @Prop({ type: Date, required: true, default: now() })
  createdAt: Date;

  @Prop({ type: Date, required: true, default: now() })
  updatedAt: Date;

  @Prop({ type: Date, default: null })
  deletedAt: Date;
}

export const TicketSchema = SchemaFactory.createForClass(Ticket);
