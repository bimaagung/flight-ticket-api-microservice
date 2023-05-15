import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { now, Date, HydratedDocument } from 'mongoose';

export type AirplaneDocument = HydratedDocument<Airplane>;

@Schema({ timestamps: true })
export class Airplane {
  @Prop({ type: String, required: true })
  id: string;

  @Prop({ required: true })
  flightCode: string;

  @Prop({ type: Date, required: true, default: now() })
  createdAt: Date;

  @Prop({ type: Date, required: true, default: now() })
  updatedAt: Date;

  @Prop({ type: Date, default: null })
  deletedAt: Date;
}

export const AirplaneSchema = SchemaFactory.createForClass(Airplane);
